const initHideMobilePhoneIconOnScroll = () => {
  const phoneIcon = document.getElementById("mobile-phone-icon")
  let oldScrollY = 0
  window.addEventListener("scroll", () => {
    if(this.scrollY < oldScrollY){
      phoneIcon.style.scale = 1
    } else{
      phoneIcon.style.scale = 0

    }
    oldScrollY = this.scrollY 
  })
}
initHideMobilePhoneIconOnScroll()