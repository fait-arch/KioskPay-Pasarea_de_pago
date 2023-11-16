import confetti from 'https://cdn.skypack.dev/canvas-confetti';

  function confettis(){
    var count = 200;
    var defaults = {};

    function fire(particleRatio, opts) {
      confetti(Object.assign({}, defaults, opts, {
        particleCount: Math.floor((count * particleRatio)+40)
      }));
    }

    fire(0.25,  {  spread: 360,  startVelocity: 55,   });
    fire(0.2,   {  spread: 360,                       });
    fire(0.35,  {  spread: 360,  decay: 0.91,  scalar: 0.8 });
    fire(0.1,   {  spread: 360,  startVelocity: 25,  decay: 0.92,  scalar: 1.2 });
    fire(0.1,   {  spread: 360,  startVelocity: 45,  });
    
    };

document.getElementById('btnConfeti').addEventListener('click',confettis);

