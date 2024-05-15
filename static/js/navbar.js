const initOpenMobileNavListener = () => {
  const HAMBURGER_ICON = ` <svg
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    stroke-width="1.5"
    stroke="currentColor"
    class="w-10 h-10"
    >
    <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
    />
    </svg>`;

  const X_ICON = `<svg xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    class="w-10 h-10">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
                </svg>
  `;
  const [getIsOpened, setIsOpened] = useState(false);
  const navMenuButton = document.getElementById("mobile-menu-button");
  navMenuButton.addEventListener("click", () => {
    if (!getIsOpened()) {
      const navDrawer = document.getElementById("nav-drawer");
      navDrawer.style.transform = "translateY(0)";
      navMenuButton.innerHTML = X_ICON;
      document.body.style.width = "100vw";
      document.body.style.height = "100vh";
      document.body.style.overflow = "hidden";
      setIsOpened(true);
    } else {
      console.log("here");
      const navDrawer = document.getElementById("nav-drawer");
      navDrawer.style.transform = "translateY(-100vh)";
      navMenuButton.innerHTML = HAMBURGER_ICON;
      document.body.style.width = "";
      document.body.style.height = "";
      document.body.style.overflow = "";
      setIsOpened(false);
    }
  });
};
initOpenMobileNavListener();
