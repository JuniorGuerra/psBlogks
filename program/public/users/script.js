/*Efecto del boton cambiante*/
const btnSwitch = document.querySelector('#switch');

btnSwitch.addEventListener('click', () => {
    btnSwitch.classList.toggle('active');
});


if (localStorage.getItem('dark-mode') === 'true') {
    btnSwitch.classList.add('active');
} else {
    btnSwitch.classList.remove('active');
}

var btn = document.getElementById("switch")
btn.addEventListener("click", style)

var fondo = document.querySelector("body")
var texto = document.getElementById("users")
var date = document.getElementById("date")
var img = document.getElementById("img")
var books = document.getElementById("books")
var p = document.getElementsByClassName("p")

var color = 0;

function style() {
    if (color == 0) {
        fondo.style.background = "#000000";
        texto.style.color = "#FFF"
        date.style.color = "#FFF"
        books.style.color = "#FFF"
        p.style.color = "#FFF"
        color = 1
    } else {
        texto.style.color = "#000"
        date.style.color = "#000"
        books.style.color = "#000"
        p.style.color = "#000"
        fondo.style.background = "#DDDDDD";

        color = 0
    }
}