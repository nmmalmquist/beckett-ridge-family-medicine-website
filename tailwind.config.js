/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js}", "./static/**/*.{html,js}"],
  safelist: ["bg-danger", "bg-success"],
  theme: {
    colors: {
      transparent: "transparent",
      current: "currentColor",
      primary: "royalblue",
      "primary-light": "#c4deff",
      secondary: "#005550",
      offwhite: "#F3F3F3",
      white: "#ffffff",
      dark: "#222831",
      lightgray: "#d3d3d3",
      success: "green",
      danger: "#8B0000",
    },
    extend: {},
  },
  plugins: [],
};
