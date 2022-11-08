import React from 'react'
import CustomButton from '../../UI/CustomButton/CustomButton'
import moment from 'moment'
import 'moment/locale/ru'

function CheckPostItem({ post }) {
    const approvePost = (e) => {
        e.preventDefault()

    }

    const rejectPost = (e) => {
        e.preventDefault()
        
    }
    return (
        <div className='mb-3 w-full p-2 flex flex-col
        rounded-md border border-solid border-[black]'>
            <h3 className='text-[28px] font-bold self-center'>{post.title}</h3>
            <h4 className='font-bold self-center'>Автор: {post.author_login} Время создания: {moment(post.created_at).format('DD.MM.YYYY hh:mm').toString()}</h4>
            <div className='w-full h-[200px] self-center'>Тут будет картинка</div>
            <h4 className='font-bold self-center'>Аннотация</h4>
            <p>{post.annotation}</p>
            <h4 className='font-bold self-center'>Содержимое</h4>
            <p>{post.text}</p>
            { post.time_start.Valid && <h4>Время начала доступности: {moment(post.time_start.Time).format('DD.MM.YYYY hh:mm').toString()}</h4> }
            { post.time_end.Valid && <h4>Время конца доступности: {moment(post.time_end.Time).format('DD.MM.YYYY hh:mm').toString()}</h4> }
            <div className='flex mt-2 self-center'>
                <span className='mr-2'><CustomButton text={'Одобрить'} callback={approvePost} /></span>
                <CustomButton text={'Отклонить'} callback={rejectPost} />
            </div>
        </div>
    )
}

export default CheckPostItem