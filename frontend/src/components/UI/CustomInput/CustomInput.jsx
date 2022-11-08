import React from 'react'

function CustomInput({ text, description, type, callback }) {
    return (
        <div className='flex flex-col my-2 '>
            <p>{description}</p>
            <input type={type} className="border border-solid border-[#457b9d] rounded-lg px-1" 
                onChange={e => callback(e.target.value)} />
        </div>
    )
}

export default CustomInput