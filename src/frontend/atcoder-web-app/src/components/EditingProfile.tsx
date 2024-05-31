import React from 'react'
import { useQueryProfile } from '../hooks/useQueryProfile'
import { useState, FormEvent, useEffect} from 'react'
import { useMutateProfile } from '../hooks/useMutateProfile'
import { useNavigate } from 'react-router-dom'
import { useQueryClient } from '@tanstack/react-query'

export const EditingProfile = () => {
    
    const {data, isLoading } = useQueryProfile()
    const [id, setId] = useState<string>('')
    const navigate = useNavigate();
    useEffect(() => {
        if (data) {
            setId(data)
        }
    }, [data])
    const queryClient = useQueryClient()

    const { updateProfileMutation } = useMutateProfile()

    const submitProfileHandler = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        updateProfileMutation.mutate(id, {
            onSuccess: () =>{
                queryClient.invalidateQueries(['tables'])
                queryClient.invalidateQueries(['submissions'])
                navigate('/user')
            }
        })
    }
    
    return (
        <div className="mb-6">
            <div className="w-full">
                <h2 className="text-xl font-bold mb-2">自分のAtcoder IDを変更</h2>
            </div>
            {isLoading ? (
                <p>Loading...</p>
            ) : (
            <form onSubmit={submitProfileHandler}>
                <input
                    className="mb-3 px-3 text-sm py-2 border border-gray-300"
                    name="id"
                    type="id"
                    autoFocus
                    placeholder="your Atcoder ID"
                    onChange={(e) => setId(e.target.value)}
                    value={id}
                />
                <button
                    type="submit"
                    className="bg-blue-500 text-white py-2 px-4 rounded"
                >
                    変更
                </button>
            </form>
            )}
        </div>
    )
}
