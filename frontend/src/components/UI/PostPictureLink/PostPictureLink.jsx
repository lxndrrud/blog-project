import React from 'react'
import { Link } from 'react-router-dom'

function PostPictureLink({ destination, filepath, altDescription }) {
    return (
        <Link to={destination} >
            <img className="h-[200px] rounded border border-solid border-[#457b9d]" 
              src={filepath} alt={altDescription} />
        </Link>
    )
}

export default PostPictureLink