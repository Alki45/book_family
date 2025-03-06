import axios from axios
const API=axios.create({
    baseURL:"http://localhost:9010",
});
export const getBooks=()=>API.get("/books/");
export const createBook=(book)=>API.post("/books/",book);