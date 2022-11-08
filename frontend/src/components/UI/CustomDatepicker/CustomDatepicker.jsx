import React from 'react'

function CustomDatepicker({ description, callback }) {
  return (
    <div className='flex flex-col my-2'>
        { description && <span>{description}</span> }
        <input
            lang="ru-RU"
            type="date"
            onChange={(e) => callback(e.target.value)} 
            className="shadow-[0_1px_3px_-2px_#9098a9] w-[210px] h-[45px]
              border border-solid border-[#457b9d] rounded-lg px-1" 
        >
        </input>
    </div>
  )
}

export default CustomDatepicker