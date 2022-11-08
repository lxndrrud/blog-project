import React from 'react'

function CustomButton({ text, callback }) {
    return (
        <button className='p-2 rounded-lg bg-[#1d3557] text-[white]' onClick={callback}>{text}</button>
    )
}

export default CustomButton