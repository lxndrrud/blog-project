import React from 'react'
import IconSVG from '../IconSVG/IconSVG'

function CustomButton({ text, callback, iconData }) {
    return (
        <span>
            <button className='p-2 rounded-lg bg-[#1d3557] text-[white]' onClick={callback}>
                { iconData && <span className='mr-2'><IconSVG data={iconData} color="#FFF" /></span> }
                {text}
            </button>
        </span> 
    )
}

export default CustomButton