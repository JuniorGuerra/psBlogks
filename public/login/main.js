function oscuro(){
const btnSwitch = document.querySelector('#switchs');
const switchs = document.getElementById("switchs");
const body = document.getElementById("body");
const form = document.getElementById("form");
const contenedor = document.getElementById("contenedor");
const correo = document.getElementById("correo");
const contra = document.getElementById("contra");
const fas = document.getElementById("fas");
const fas2 = document.getElementById("fas2");
const inicio = document.getElementById("inicio");
const olvido = document.getElementById("olvido");
const link = document.getElementById("link");
const no_cuenta = document.getElementById("no_cuenta");


    btnSwitch.addEventListener('click', () => {
    document.body.classList.toggle('dark');
    btnSwitch.classList.toggle('active');

    body.classList.toggle("body_oscuro");
    form.classList.toggle("formulario_oscuro");
    contenedor.classList.toggle("contenedor_oscuro");
    correo.classList.toggle("correo_oscuro");
    fas.classList.toggle("fas_oscuro");
    fas2.classList.toggle("fas_oscuro");
    inicio.classList.toggle("button_oscuro");
    olvido.classList.toggle("text_oscuro");
    link.classList.toggle("link_oscuro");
    contra.classList.toggle("contra_oscuro");
    no_cuenta.classList.toggle("text_oscuro");
})
}