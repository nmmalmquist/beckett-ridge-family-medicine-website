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
      offwhite: "#F3F3F3",
      white: "#ffffff",
      dark: "#222831",
      lightgray: "#d3d3d3",
      success: "green",
      danger: "red",
    },
    extend: {},
  },
  plugins: [],
};
