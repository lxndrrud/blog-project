import React from 'react'
import { Link } from 'react-router-dom'

function HeaderLink({destination, text }) {
    return (
       <Link to={destination} className="mr-3 p-1 rounded-lg">{text}</Link>
    )
}

export default HeaderLink