import {
  ArrowRightOnRectangleIcon,
  ShieldCheckIcon,
} from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { SubmissionList } from './SubmissionList'
import { useQueryClient } from '@tanstack/react-query'


export const User = () => {
  const queryClient = useQueryClient()
  const { logoutMutation } = useMutateAuth()
  const logout = async () => {
    await logoutMutation.mutateAsync()
    queryClient.removeQueries(['submissions']) // キャッシュをクリア
  }
  return (
    <div>
      <ArrowRightOnRectangleIcon
        onClick={logout}
        className="h-6 w-6 my-6 text-blue-500 cursor-pointer"
      />
      <SubmissionList/>
    </div>
  )
}
