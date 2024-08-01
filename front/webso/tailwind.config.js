/** @type {import('tailwindcss').Config} */
export default {
    content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
    theme: {
        extend: {
            keyframes: {
                pulsik: {
                    "50%": { opacity: 0.5 },
                },
            },
            animation: {
                pulsik: "pulsik 0.5s ease-in-out infinite",
            },
        },
    },
    plugins: [],
};