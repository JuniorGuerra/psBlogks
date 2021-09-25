
/*Efecto del boton cambiante*/
const btnSwitch = document.querySelector('#switch');

btnSwitch.addEventListener('click', () => {
	btnSwitch.classList.toggle('active');
});


if(localStorage.getItem('dark-mode') === 'true'){
	btnSwitch.classList.add('active');
} else {
	btnSwitch.classList.remove('active');
}




var btn = document.getElementById("switch")
btn.addEventListener("click", style)

var fondo = document.querySelector("body")
var title = document.getElementById("title")
var text = document.getElementById("text")
var b_login = document.getElementById("b_login")
var b_register = document.getElementById("b_register")
var llanada = document.getElementById("llamada")



var color = 0;
function style() {
    if (color == 0) {
        fondo.style.background = "#000000";
        title.style.color = "#00E1D8";
        title.style.background = "-webkit-linear-gradient(yellow,red);"
        text.style.color = "#DDDDDD"
        b_login.style.color = "#FFFFFF"
        b_register.style.color = "#C80F53"
        b_register.style.background = "#DDDDDD"
        llanada.style.background = "#C80F53"
        color = 1
    } else {
        fondo.style.background = "#DDDDDD";
        title.style.color = "black";
        text.style.color = "#000"
        llanada.style.background = "#000"
        b_login.style.color = "#000"
        b_register.style.color = "#DDDDDD"
        b_register.style.background = "#000"
        color = 0
    }
}





main()
var contador = 1;

function main () {
    var menu_bar = document.getElementById('menu_bar')
    var nav = document.getElementById('nav')
	menu_bar.addEventListener("click",function(){
		if (contador == 1) {
			nav.style.left = "0"
			contador = 0;
        llanada.innerHTML = "Menu"
		} else {
			contador = 1;
            nav.style.left="-100%"
            llanada.innerHTML = ""
		}
	});

	// Mostramos y ocultamos submenus

}



/*
Colores de la pagina:

Negro: 000000
Cian: 00E1D8
Fusia: C80F53
Grisesito: DDDDDD
*/
