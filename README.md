mutation CreateUser{
createUser(input:{username:"earl", password:"passwd"})
}

---

mutation UserLogin{
login(input:{
username:"earl",
password:"passwd"
})
}

---

mutation CreateLink{
createLink(input: {title: "user1Title", address: "user1Address"})
{
user{name}
}
}

Headers
{
"Authorization":"TOKEN"
}

---

query Link{
links{
id
user {
id
name
}
title
address
}
}