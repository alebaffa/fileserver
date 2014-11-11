package main

import (
    "math/rand"
    "github.com/go-martini/martini"
    "github.com/render"
    "github.com/binding"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Team struct {
    Team  string `form:"teamname"`
    Members []string `form:"teammate[]"`
}

/* 
   the function returns a martini.Handler which is called on each request. We simply clone 
   the session for each request and close it when the request is complete. The call to c.Map 
   maps an instance of *mgo.Database to the request context. Then *mgo.Database
   is injected into each handler function.
*/
func DB() martini.Handler {
  session, err := mgo.Dial("mongodb://localhost") // mongodb://localhost
  if err != nil {
    panic(err)
  }
 
  return func(c martini.Context) {
    s := session.Clone()
    c.Map(s.DB("sheriff")) // local
    defer s.Close()
    c.Next()
  }
}

// function to return an array of all team members from mongodb
func getTeam(db *mgo.Database) Team {
  var team Team
  db.C("testData").Find(bson.M{"team": "DD2"}).One(&team)
  return team
} 

func main() {
    m := martini.Classic()
    m.Use(render.Renderer(render.Options {
    Layout: "layout"}))
    // use the Mongo middleware
    m.Use(DB())
    //http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
    m.Get("/", func(r render.Render, db *mgo.Database) {
        r.HTML(200, "home", nil)
    }) 

/* 
    create a new Team the form submission. Contains some martini magic. The call 
    to binding.Form(Team{}) parses out form data when the request comes in. 
    It binds the data to the struct, maps it to the request context  and
    injects into our next handler function to insert into Mongodb.
 */   
    m.Post("/subscribe", binding.Form(Team{}), func(team Team, r render.Render, db *mgo.Database) {
        db.C("testData").Insert(team)
        r.HTML(200, "home", nil)
    })

    m.Post("/assign", func(r render.Render, db *mgo.Database){
        team := getTeam(db)
        r.HTML(200, "winner", team.Members[rand.Intn(len(team.Members))])
    })

    m.Run()
}