import React from 'react'
import CustomButton from '../../UI/CustomButton/CustomButton'
import { useNavigate } from 'react-router-dom'


function DialogForm({ question, submitCallback, declineDestination }) {
    const navigate = useNavigate()
    return (
        <div className='w-full flex flex-col'>
            <span className='self-center'>{question}</span>
            <div className='flex self-center'>
                <span className='mt-2 mr-2'><CustomButton text={'Подтвердить'} callback={(e) => submitCallback(e.target.value)} /></span>
                <span className='mt-2 mr-2'><CustomButton text={'Назад'} callback={(e) => navigate(declineDestination)} /></span>
            </div>
        </div>
    )
}

export default DialogForm