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
var img = document.getElementById("img")
var color = 0;

function style() {
    if (color == 0) {
        fondo.style.background = "#000000";
        texto.style.color = "#FFF"
        color = 1
    } else {
        texto.style.color = "#000"
        fondo.style.background = "#DDDDDD";

        color = 0
    }
}