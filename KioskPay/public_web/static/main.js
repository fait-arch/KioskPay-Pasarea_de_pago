document.addEventListener('DOMContentLoaded', function() {
    // Acción abrir la ventana modal
    var modal = document.getElementById('shoppingCartModal');
    var btns = document.querySelectorAll('.add-to-cart-btn');
    btns.forEach(btn => {
        btn.onclick = function() {
            modal.style.display = "block";
        }
    });

    // Acción cerrar la ventana modal
    var span = document.getElementsByClassName("close")[0];
    if(span) {
        span.onclick = function() {
            modal.style.display = "none";
            reiniciarDatos();
        }
    }

    // Función para reiniciar los datos del carrito y el formulario
    function reiniciarDatos() {
        var billingForm = document.getElementById('billingForm');
        if(billingForm) billingForm.reset();

        var camposCarrito = ['itemName', 'itemDescription', 'itemPrice', 'discountAmount', 'totalPrice', 'totalPriceDisplay'];
        camposCarrito.forEach(function(campo) {
            var elemento = document.getElementById(campo);
            if(elemento) elemento.textContent = campo === 'totalPriceDisplay' ? "PRECIO FINAL: $0.00" : "$0.00";
        });

        var cantidadInputs = document.querySelectorAll('#cartTable input[type=number]');
        cantidadInputs.forEach(input => {
            input.value = 1;
        });

        var promoCodeInput = document.querySelector('.promo-code input');
        if(promoCodeInput) promoCodeInput.value = '';

        actualizarPrecioTotal();
    }

    // Función para cambiar entre pestañas
    function openTab(evt, tabName) {
        var tabcontent = document.getElementsByClassName("tabcontent");
        for (var i = 0; i < tabcontent.length; i++) {
            tabcontent[i].style.display = "none";
        }

        var tablinks = document.getElementsByClassName("tablink");
        for (var i = 0; i < tablinks.length; i++) {
            tablinks[i].className = tablinks[i].className.replace(" active", "");
        }

        document.getElementById(tabName).style.display = "block";
        evt.currentTarget.className += " active";
    }

    var productos = [
        {
          id: 1,
          nombre: "Producto 1 - Cereal",
          descripcion: "Descripción breve del producto...",
          precio: 4.99
        },

        {
          id: 2,
          nombre: "Producto 2 - Bebida",
          descripcion: "Descripción breve del producto...",
          precio: 3.50
        },

        {
          id: 3,
          nombre: "Producto 3 - Carnes",
          descripcion: "Descripción breve del producto...",
          precio: 5.25
        },

        {
          id: 4,
          nombre: "Producto 4 - limpieza",
          descripcion: "Descripción breve del producto...",
          precio: 6.00
        },
        
      ];

      // Función para actualizar el precio total de un producto
      function actualizarPrecioTotal() {
        var cantidad = document.querySelector('#cartTable input[type=number]').value;
        var precioUnitario = parseFloat(document.getElementById('itemPrice').textContent.replace('$', ''));
        var precioTotalSinDescuento = cantidad * precioUnitario;
        
        // Aplicar descuento si es necesario y calcular el precio total final
        var descuento = aplicarDescuento(precioTotalSinDescuento);
        var precioTotalConDescuento = precioTotalSinDescuento - descuento;

        document.getElementById('totalPrice').textContent = '$' + precioTotalConDescuento.toFixed(2);
        document.getElementById('totalPriceDisplay').textContent = 'PRECIO FINAL: $' + precioTotalConDescuento.toFixed(2);
      }
      // Función para aplicar el descuento
      function aplicarDescuento(precioTotal) {
        var codigoPromo = document.querySelector('.promo-code input').value;
        var descuento = 0;

        if (codigoPromo === '12345') {
          descuento = precioTotal * 0.25; // 25% de descuento
        }
        if (codigoPromo === 'JuanDaCode') {
          descuento = precioTotal * 0; // 0% de descuento
        }
        if (codigoPromo === 'JuanDaCode') {
          descuento = precioTotal * 0; // 0% de descuento
        }

        document.getElementById('discountAmount').textContent = '-$' + descuento.toFixed(2);
        return descuento;
      }
      // Evento al botón de aplicar el código promocional
      document.querySelector('.apply-promo-btn').addEventListener('click', function() {
        actualizarPrecioTotal();
      });

    var applyPromoBtn = document.querySelector('.apply-promo-btn');
    if(applyPromoBtn) {
        applyPromoBtn.addEventListener('click', function() {
            actualizarPrecioTotal();
        });
    }

    var cartTableInput = document.querySelector('#cartTable input[type=number]');
    if(cartTableInput) {
        cartTableInput.addEventListener('change', function() {
            actualizarPrecioTotal();
        });
    }

    function agregarAlCarrito(idProducto) {
        var producto = productos.find(p => p.id === idProducto);
        if (producto) {
          // Asigna los valores a los campos del carrito
          document.getElementById('itemName').textContent = producto.nombre;
          document.getElementById('itemDescription').textContent = producto.descripcion;
          document.getElementById('itemPrice').textContent = "$" + producto.precio.toFixed(2);
          // Actualiza el precio total
          actualizarPrecioTotal();
          openTab(event, 'Cart');
        } 
      }

    var botonesAgregar = document.querySelectorAll('.add-to-cart-btn');
    botonesAgregar.forEach(btn => {
        btn.onclick = function() {
            var productoId = parseInt(btn.getAttribute('data-producto-id'));
            agregarAlCarrito(productoId);
            modal.style.display = "block";
        };
    });

    var nextStepButton = document.getElementById('nextStepButton');
    if(nextStepButton) {
        nextStepButton.addEventListener('click', function() {
            var nextTab = this.getAttribute('data-next-tab');
            openTab(event, nextTab);
        });
    }

    // Más funciones según sea necesario

});
