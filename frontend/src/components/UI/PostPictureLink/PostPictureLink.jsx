import React from 'react'
import { Link } from 'react-router-dom'

function PostPictureLink({ destination, filepath, altDescription}) {
    return (
        <Link to={destination} className="w-[max-content] mx-auto" >
            <img className="mx-2 h-[200px] rounded border border-solid border-[#457b9d]" 
              src={`/storage/${filepath}`} alt={altDescription} />
        </Link>
    )
}

export default PostPictureLink