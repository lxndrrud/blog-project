import React from 'react'

function PostPictureImage({ filepath, altDescription }) {
    return (
        <img className="mx-2 h-[200px] self-center rounded border border-solid border-[#457b9d]" 
            src={`/storage/${filepath}`} alt={altDescription} />
    )
}

export default PostPictureImage