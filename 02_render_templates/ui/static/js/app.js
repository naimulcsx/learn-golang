document.addEventListener("DOMContentLoaded", () => {
  console.log("App initialized");

  // Add any interactive features here
  const navLinks = document.querySelectorAll("nav a");
  navLinks.forEach((link) => {
    link.addEventListener("click", (e) => {
      console.log(`Navigating to: ${e.target.href}`);
    });
  });
});
