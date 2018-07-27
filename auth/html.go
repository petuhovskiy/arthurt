package main

const loginHTML = `
<!DOCTYPE html>
<html>
Please login!
<br>
<form action="/login" method="POST">
<input type="text" name="user" placeholder="Username"/>
<br>
<input type="password" name="password" placeholder="Password"/>
<button type="submit">Login</button>
</form>
</html>
`

const registerHTML = `
<!DOCTYPE html>
<html>
Register
<br>
<form action="/register" method="POST">
<input type="text" name="user" placeholder="Username"/>
<br>
<input type="password" name="password" placeholder="Password"/>
<button type="submit">Register</button>
</form>
</html>
`
