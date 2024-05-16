/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js}"],
  theme: {
    colors: {
      transparent: "transparent",
      current: "currentColor",
      primary: "royalblue",
      "primary-light": "#c4deff",
      secondary: "#005550",
      offwhite: "#EEEEEE",
      white: "#ffffff",
      dark: "#222831",
      lightgray: "#d3d3d3",
    },
    extend: {},
  },
  plugins: [],
};
