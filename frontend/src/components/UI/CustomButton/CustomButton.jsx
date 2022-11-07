import React from 'react'

function CustomButton({ text, callback }) {
    return (
        <button className='p-2 rounded-lg bg-[lightgreen]' onClick={callback}>{text}</button>
    )
}

export default CustomButton