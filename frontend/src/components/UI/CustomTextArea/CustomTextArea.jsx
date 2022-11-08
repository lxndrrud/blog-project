import React from 'react'

function CustomTextArea({ description, text, callback, rows=5, cols=10 }) {
    return (
        <div className='flex flex-col my-2 '>
            { description && <p>{description}</p>}
            <textarea 
                onChange={(e) => callback(e.target.value)}
                rows={rows}
                cols={cols}
                value={text}
                className="border border-solid border-[#457b9d] rounded-lg px-1"
            ></textarea>
        </div>
    )
}

export default CustomTextArea