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
var text = document.getElementById("text")

var color = 0;

function style() {
    if (color == 0) {
        fondo.style.background = "#000000";
        text.style.color = "#FFF"
        color = 1
    } else {
        text.style.color = "#000"
        fondo.style.background = "#DDDDDD";
        color = 0
    }
}