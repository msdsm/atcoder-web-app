import { FC, memo } from 'react'
import { Rival } from '../types'
import { TrashIcon } from '@heroicons/react/24/solid'
import { useMutateRival } from '../hooks/useMutateRival'

interface EditingRivalItemProps extends Rival {
    refetch: () => void;
}

export const EditingRivalItem: FC<EditingRivalItemProps> = ({
    id,
    atcoder_id,
    refetch,
}) => {
    const { deleteRivalMutation } = useMutateRival()
    const deleteRivalHandler = () => {
        deleteRivalMutation.mutate(id, {
            onSuccess: () => {
                refetch()
            },
        })
    }
    return (
        <tr className="even:bg-gray-100">
            <td className="px-6 py-4">{atcoder_id}</td>
            <td className="px-6 py-4">
                <TrashIcon
                className="h-5 w-5 text-blue-500 cursor-pointer"
                onClick={deleteRivalHandler}
                />
            </td>
        </tr>
    )
}