import React from 'react'
import { useMutateRival } from '../hooks/useMutateRival'
import { useQueryRivals } from '../hooks/useQueryRivals'
import useStore from '../store'
import {
    ArrowRightOnRectangleIcon,
    ShieldCheckIcon,
} from '@heroicons/react/24/solid'
import { Link } from 'react-router-dom';
import { EditingRivalList } from './EditingRivalList'
import { EditingProfile } from './EditingProfile'
import { useNavigate } from 'react-router-dom'
import { useQueryClient } from '@tanstack/react-query'

export const Editing = () => {
    const queryClient = useQueryClient()
    const navigate = useNavigate();
    const navigateUser = () => {
      queryClient.invalidateQueries(['tables'])
      queryClient.invalidateQueries(['submissions'])
      navigate('/user')
    }
    return (
      <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
         <div className="flex items-center my-3">
          <ShieldCheckIcon className="h-8 w-8 mr-3 text-indigo-500 cursor-pointer"/>
          <span className="text-center text-3xl font-extrabold">
            Atcoder Rival App
          </span>
        </div>
        <div className="my-10"></div>
        <EditingProfile/>
        <div className="my-10"></div>
        <EditingRivalList/>
        <div className="my-10"></div>
        <span className="ml-2 text-blue-500" onClick={navigateUser}>ライバルユーザーリスト編集完了</span>
        <div className="my-10"></div>
      </div>
    )
}
