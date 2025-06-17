// Preloader
window.addEventListener("load", function () {
  setTimeout(function () {
    const preloader = document.querySelector(".preloader");
    preloader.style.opacity = "0";
    setTimeout(function () {
      preloader.style.display = "none";
    }, 500);
  }, 1000); // 1 second loading animation
});

// Initialize AOS
AOS.init({
  duration: 1000,
  once: true,
});

// Typed.js initialization
document.addEventListener("DOMContentLoaded", function () {
  new Typed("#typed-text", {
    strings: [
      "Python &amp; Java Developer",
      "Web &amp; Software Enthusiast",
      "Problem Solver",
      "Continuous Learner",
    ],
    typeSpeed: 50,
    backSpeed: 30,
    backDelay: 2000,
    loop: true,
  });
});

// Particles.js config
particlesJS("particles-js", {
  particles: {
    number: { value: 80, density: { enable: true, value_area: 800 } },
    color: { value: "#4a6bff" },
    shape: { type: "circle" },
    opacity: { value: 0.5, random: false },
    size: { value: 3, random: true },
    line_linked: {
      enable: true,
      distance: 150,
      color: "#4a6bff",
      opacity: 0.4,
      width: 1,
    },
    move: {
      enable: true,
      speed: 2,
      direction: "none",
      random: false,
      straight: false,
      out_mode: "out",
      bounce: false,
    },
  },
  interactivity: {
    detect_on: "canvas",
    events: {
      onhover: { enable: true, mode: "grab" },
      onclick: { enable: true, mode: "push" },
      resize: true,
    },
  },
});

// Contact form submission
document.addEventListener("DOMContentLoaded", function () {
  const contactForm = document.getElementById("contactForm");
  if (contactForm) {
    contactForm.addEventListener("submit", function (e) {
      const submitBtn = this.querySelector(".form-submit");
      submitBtn.innerHTML =
        '<i class="fas fa-spinner fa-spin me-2"></i> Sending...';
      submitBtn.disabled = true;
    });
  }
});

// Navbar scroll effect
window.addEventListener("scroll", function () {
  const navbar = document.querySelector(".navbar");
  if (window.scrollY > 50) {
    navbar.style.padding = "10px 0";
    navbar.style.boxShadow = "0 2px 10px rgba(0, 0, 0, 0.1)";
  } else {
    navbar.style.padding = "15px 0";
    navbar.style.boxShadow = "none";
  }
});

let ticking = false;

function updateActiveNavigation() {
  const sections = document.querySelectorAll("section[id]");
  const navLinks = document.querySelectorAll(".navbar-nav .nav-link");
  const navbarHeight = document.querySelector(".navbar").offsetHeight;
  const scrollPosition = window.scrollY + navbarHeight + 50;

  let currentSection = "";

  //find current section
  sections.forEach((section) => {
    const sectionTop = section.offsetTop;
    const sectionHeight = section.offsetHeight;
    const sectionId = section.getAttribute("id");

    if (
      scrollPosition >= sectionTop &&
      scrollPosition < sectionTop + sectionHeight
    ) {
      currentSection = sectionId;
    }
  });

  if (window.scrollY < 100) {
    currentSection = "hero";
  }

   // Map section IDs to navigation link text
  const sectionToNavText = {
    hero: "Home",
    about: "About",
    projects: "Projects",
    certificates: "Certificates",
    contact: "Contact",
  };

  navLinks.forEach(link =>{
    const linkText = link.textContent.trim();

    link.classList.remove('active')

    if (sectionToNavText[currentSection] === linkText){
        link.classList.add('active')
    }
  });

  ticking = false
}

//Optimized scroll handler
function onScroll(){
    if(!ticking){
        requestAnimationFrame(updateActiveNavigation);
        ticking = true;
    }
}

// Scroll to top button logic
const scrollToTopBtn = document.querySelector(".scroll-to-top");

function toggleScrollToTopButton() {
  if (window.pageYOffset > 300) {
    scrollToTopBtn.classList.add("active");
  } else {
    scrollToTopBtn.classList.remove("active");
  }
}

// Smooth scrolling setup
document.addEventListener('DOMContentLoaded', function(){
    const navLinks = document.querySelectorAll('.navbar-nav .nav-link');

    // Add click handlers to nav links for smooth scrolling    
    navLinks.forEach(link =>{
        link.addEventListener('click', function(e){
            const linkText = this.textContent.trim();

            //Map nav text to section IDs
            const navTextToSection = {
                'Home': 'hero',
                'About': 'about', 
                'Projects': 'projects',
                'Certificates': 'certificates',
                'Contact': 'contact'
            };

            const targetSectionId = navTextToSection[linkText];
            const targetSection = document.getElementById(targetSectionId);

            // If we found a matching section on the current page, scroll to it
            if(targetSection){
                e.preventDefault();
            }

            //Close mobile menu if open
            const navbarCollapse = document.querySelector('.navbar-collapse');

            if(navbarCollapse && navbarCollapse.classList.contains('show')){
                const toggleBtn = document.querySelector('.navbar-toggler');
                if(toggleBtn) toggleBtn.click();
            }

            const navbarHeight = document.querySelector('.navbar').offsetHeight;
            const targetPosition = targetSection.offsetTop - navbarHeight;

            window.scrollTo({
                top: targetPosition,
                behavior: 'smooth'
            });

            //Update active link immediately
            setTimeout(() =>{
                updateActiveNavigation();
            },100)
        })
    })

    // Handle scroll to top button
  if (scrollToTopBtn) {
    scrollToTopBtn.addEventListener("click", () => {
      window.scrollTo({ top: 0, behavior: "smooth" });
    });
  }

   // Initial setup
  updateActiveNavigation();
   toggleScrollToTopButton();
})



// Attach  scroll listener
window.addEventListener('scroll', function(){
    onScroll();
    toggleScrollToTopButton();
}, {passive : true });