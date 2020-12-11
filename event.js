use admin
db.createUser(
  {
    user: "myUserAdmin",
    pwd: "926443",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ]
  }
)


docker run --name bluebell_mongo -d -p 27017:27017 -v mongoData:/data/db mongo


db.events.insert({
    "event_id": 123456789,
    "event_name": "Philadelphia Eagles vs. Seattle Seahawks",
    "event_date_and_time": 1606767300000,
    "cast": [
        "Philadelphia Eagles",
        "Seattle Seahawks"
    ],
    "event_address": {
        "street": "Lincoln Financial Field",
        "city": "Philadelphia",
        "State": "PA",
        "country": "United States",
        "zipcode": "08873"
    },
    "event_type": "sports",
    "ticket_type": [
        {
            "level": "vip",
            "price": 500,
            "total_volumn": 10
        },
        {
            "level": "student",
            "price": 50,
            "total_volumn": 10
        },
        {
            "level": "normal",
            "price": 150,
            "total_volumn": 100
        }
    ],
    "ticketing_start_time": 1606680900000,
    "event_description": "Last chance! Come to see!"
})