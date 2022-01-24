module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        "active-link": "#ED1980",
        "post-border": "#424242",
      },
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
