import React from "react";
import {
    Link
} from "react-router-dom";

function Home() {
    console.log("Home");
    return (
        <>
        <p>Home page</p>
        <ul>
            <li>
                <Link to="/login">Login</Link>
            </li>
            <li>
                <Link to="/signup">Signup</Link>
            </li>
            <li>
                <Link to="/dashboard">Dashboard</Link>
            </li>
            <li>
                <Link to="/">Home</Link>
            </li>
        </ul>
        </>
    )
}

export default Home;