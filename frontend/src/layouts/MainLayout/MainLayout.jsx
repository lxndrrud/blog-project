import React from 'react'
import Header from '../../components/UI/Header/Header'

function MainLayout({ title, children }) {
    return (
        <div className='px-1 sm:px-0 sm:w-[60%] flex flex-col mx-auto'>
            <Header />
            <span className='mt-1 text-[26px]'>{title}</span>
            <div className='w-full flex flex-col'>
                { children }
            </div>
        </div>
    )
}

export default MainLayout