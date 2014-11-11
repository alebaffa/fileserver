package main

import (
    "math/rand"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/martini-contrib/binding"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "fmt"
    "os"
)

type Team struct {
    Team  string `form:"teamname"`
    Members []string `form:"teammate[]"`
}

type TeamName struct {
    Name string `form:"team_name"`
}

/* 
   the function returns a martini.Handler which is called on each request. We simply clone 
   the session for each request and close it when the request is complete. The call to c.Map 
   maps an instance of *mgo.Database to the request context. Then *mgo.Database
   is injected into each handler function.
*/
func DB() martini.Handler {
  session, err := mgo.Dial(os.Getenv("MONGOLAB_URI")) // mongodb://localhost
  if err != nil {
    panic(err)
  }
 
  return func(c martini.Context) {
    s := session.Clone()
    c.Map(s.DB(os.Getenv("MONGO_DB"))) // local
    defer s.Close()
    c.Next()
  }
}

// function to return an array of all team members from mongodb
func getTeam(db *mgo.Database, name string) Team {
  var team Team
  db.C("teams").Find(bson.M{"team": name}).One(&team)
  return team
} 

// function to return an array of all team members from mongodb
func All(db *mgo.Database) []Team {
  var teams []Team
  db.C("teams").Find(bson.M{}).All(&teams)
  return teams
}

var global_team_name string
var global_teams []Team

func insertMongo(team Team, db *mgo.Database) {
    members_valid := make([]string, 1)
    for i := 0; i < len(team.Members); i++ {
        if team.Members[i] != "" {
            members_valid = append(members_valid, team.Members[i])
        }
    }
    doc := Team{Team: team.Team, Members: members_valid}
    db.C("teams").Insert(doc)
}

func main() {
    m := martini.Classic()
    m.Use(render.Renderer())
    // use the Mongo middleware
    m.Use(DB())
    //http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
    m.Get("/", func(r render.Render, db *mgo.Database) {
        global_team_name = ""
        global_teams = All(db)
        r.HTML(200, "home", global_teams)
    }) 

/* 
    create a new Team the form submission. Contains some martini magic. The call 
    to binding.Form(Team{}) parses out form data when the request comes in. 
    It binds the data to the struct, maps it to the request context  and
    injects into our next handler function to insert into Mongodb.
 */   
    m.Post("/subscribe", binding.Form(Team{}), func(team Team, r render.Render, db *mgo.Database) {
        insertMongo(team, db)
        global_teams = append(global_teams, team)
        r.HTML(200, "home", global_teams)
    })

    m.Post("/set_team", binding.Form(TeamName{}), func(name TeamName, r render.Render, db *mgo.Database) {
        global_team_name = name.Name
        fmt.Println("set global variable to: ", global_team_name)
        r.HTML(200, "home", global_teams)
    })

    m.Post("/assign", func(r render.Render, db *mgo.Database){
        if global_team_name == "" {
            r.HTML(200, "home", global_teams)
        } else {
            team := getTeam(db, global_team_name)
            fmt.Println("found team: ", team)
            r.HTML(200, "winner", team.Members[rand.Intn(len(team.Members))])    
        }
        
    })

    m.Run()
}