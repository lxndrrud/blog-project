import React from 'react'
import { Link } from 'react-router-dom'
import IconSVG from '../IconSVG/IconSVG'

function HeaderLink({ destination, text, iconData }) {
    return (
        <Link to={destination} className="mr-3 p-1 rounded-lg">
            { iconData && <IconSVG data={iconData} /> }
            {text}
        </Link>
    )
}

export default HeaderLink