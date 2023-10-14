const { query } = require('express')
const express = require('express')
const app = express()
const port = 3000
const config = {
    'host': 'mysql',
    'user': 'root',
    'password': 'root',
    'database': 'nodedb'
}

const mysql = require('mysql')
const connection = mysql.createConnection(config)

const migration = {
    run: (conecction) => { conecction.query(`create table if not exists people(id int not null auto_increment, name varchar(255), primary key (id));`) }
}

migration.run(connection)

connection.query(`INSERT INTO people(name) values('Greg0x46')`)

let user

connection.query("SELECT * FROM people ORDER BY id DESC LIMIT 1;", function (error, results, fields) {
    if(error){
        console.log(error)
    }

    user = results[0].name
})

connection.end()

app.get("/", (req, res) => {
    let html = "<h1>Full Cycle</h1><br><h3>Hello " + user + " </h3>"
    res.send(html)
})

app.listen(port, () => {
    console.log("Running on port " + port)
})
