import {
  ArrowRightOnRectangleIcon,
  ShieldCheckIcon,
} from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { SubmissionList } from './SubmissionList'
import { TableList } from './TableList'
import { useQueryClient } from '@tanstack/react-query'
import { Link } from 'react-router-dom';

export const User = () => {
  const queryClient = useQueryClient()
  const { logoutMutation } = useMutateAuth()
  const logout = async () => {
    await logoutMutation.mutateAsync()
    queryClient.removeQueries(['submissions']) // キャッシュをクリア
  }
  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
       <div className="flex items-center my-3">
        <ShieldCheckIcon className="h-8 w-8 mr-3 text-indigo-500 cursor-pointer"/>
        <span className="text-center text-3xl font-extrabold">
          Atcoder Rival App
        </span>
      </div>
      <div className="flex items-center">
      <ArrowRightOnRectangleIcon
        onClick={logout}
        className="h-6 w-6 my-6 text-blue-500 cursor-pointer"
      />
      <span className="ml-2 text-blue-500" onClick={logout}>ログアウト</span>
      </div>
      <TableList />
      <Link to="/editing" className="m1-2 text-blue-500">
        ユーザーリスト編集
      </Link>
      <div className="my-10"></div>
      <SubmissionList/>
    </div>
  )
}
