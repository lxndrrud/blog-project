import React from 'react'
import Header from '../../components/UI/Header/Header'

function MainLayout({ title, children }) {
    return (
        <div className='min-h-[100vh] sm:w-[60%] flex flex-col mx-auto shadow-2xl bg-[#a8dadc]'>
            <Header />
            <div className='px-1 sm:px-5 w-full flex flex-col '>
                <span className='mt-5 mb-3 text-[32px]'>{title}</span>
                <div className='w-full flex flex-col'>
                    { children }
                </div>
            </div>
        </div>
    )
}

export default MainLayout