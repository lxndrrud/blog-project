import React from 'react'
import { Link } from 'react-router-dom'

function CustomLink({ text, destination }) {
    return (
        <Link to={destination} className="p-2 rounded-lg bg-[#1d3557] text-[white]" >{text}</Link>
    )
}

export default CustomLink