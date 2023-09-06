// Obtener una referencia al botón y al div de resultado
const miBoton = document.getElementById("miBoton");
const resultado = document.getElementById("resultado");

// Agregar un evento de clic al botón
miBoton.addEventListener("click", function () {
    resultado.textContent = "¡Hiciste clic en el botón!";
});