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
var title = document.getElementById("user")
var resumen = document.getElementById("resumen")
var email = document.getElementById("email")
var b_login = document.getElementById("b_login")
var category = document.getElementById("category")
var b_register = document.getElementById("b_register")
var llanada = document.getElementById("llamada")



var color = 0;
function style() {
    if (color == 0) {
        fondo.style.background = "#000000";
        title.style.color = "#00E1D8";
        resumen.style.color = "#DDDDDD"
        email.style.color = "#DDDDDD"
        category.style.color="#DDDDDD"
        color = 1
    } else {
        fondo.style.background = "#DDDDDD";
        title.style.color = "black";
        color = 0
        resumen.style.color = "black"
        email.style.color = "black"
        category.style.color="black"
    }
}
