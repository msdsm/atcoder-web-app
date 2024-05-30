import React from 'react'
import { useMutateRival } from '../hooks/useMutateRival'
import { useQueryRivals } from '../hooks/useQueryRivals'
import useStore from '../store'
import {
    ArrowRightOnRectangleIcon,
    ShieldCheckIcon,
} from '@heroicons/react/24/solid'
import { Link } from 'react-router-dom';

export const Editing = () => {
    const {data, isLoading } = useQueryRivals()
    const { editedRival } = useStore()
    const updateTable = useStore((state) => state.updateEditedRival)
    const {createRivalMutation, updateRivalMutation, deleteRivalMutation} = useMutateRival()
    console.log(data)
    return (
      <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
         <div className="flex items-center my-3">
          <ShieldCheckIcon className="h-8 w-8 mr-3 text-indigo-500 cursor-pointer"/>
          <span className="text-center text-3xl font-extrabold">
            Atcoder Rival App
          </span>
        </div>
        <Link to="/user" className="m1-2 text-blue-500">
          ユーザーリスト編集完了
        </Link>
        <div className="my-10"></div>
      </div>
    )
}
