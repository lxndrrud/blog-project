import React from 'react'
import CreatePostForm from '../components/Forms/CreatePostForm/CreatePostForm'
import MainLayout from '../layouts/MainLayout/MainLayout'

function CreatePostPage() {
    return (
        <MainLayout title={'Написание шедевра'} >
            <CreatePostForm />
        </MainLayout>
    )
}

export default CreatePostPage