
    function oscuro(){
        const btnSwitch = document.querySelector('#switchs');
        const switchs = document.getElementById("switchs");
        const body = document.getElementById("body");
        const form = document.getElementById("form");
        const contenedor = document.getElementById("i-contenedor");
        const text = document.getElementById("text");
        const boton = document.getElementById("boton");
        const link = document.getElementById("link");
        const h1 = document.getElementById("h");
        const texto = document.getElementById("texto");
        const cancelar = document.getElementById("cancelar");
        
        
            btnSwitch.addEventListener('click', () => {
            document.body.classList.toggle('dark');
            btnSwitch.classList.toggle('active');
            body.classList.toggle("body_oscuro");
            form.classList.toggle("formulario_oscuro");
            text.classList.toggle("text_oscuro");
            boton.classList.toggle("button_oscuro");
            h1.classList.toggle("h_oscuro");
            texto.classList.toggle("texto_oscuro");
            link.classList.toggle("link_oscuro");
            cancelar.classList.toggle("cancelar_oscuro");
        })
        }