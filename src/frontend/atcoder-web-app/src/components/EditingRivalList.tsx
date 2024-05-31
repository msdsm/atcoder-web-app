import React from 'react'
import { useQueryRivals } from '../hooks/useQueryRivals'
import { useState, FormEvent, useEffect} from 'react'
import { useMutateRival } from '../hooks/useMutateRival'
import { useNavigate } from 'react-router-dom'
import { useQueryClient } from '@tanstack/react-query'
import { Rival } from '../types'
import { EditingRivalItem } from './EditingRivalItem'

export const EditingRivalList = () => {
    const {data, isLoading, refetch} = useQueryRivals()
    const navigate = useNavigate();
    const [rivals, setRivals] = useState<Rival[]>()
    useEffect(() => {
        if (data) {
            setRivals(data)
        }
    }, [data])
    const [newId, setNewId] = useState<string>('')
    const queryClient = useQueryClient()

    const { createRivalMutation, deleteRivalMutation } = useMutateRival()
    
    const submitRivalHandler = async (e: FormEvent<HTMLFormElement>) => {
        console.log("submitRivalHandler")
        console.log(newId)
        e.preventDefault()
        createRivalMutation.mutate({
            rival: newId
        },
        {
            onSuccess: () =>{
                refetch()
            }
        })
    }
    return (
        <div className="mb-6">
            <div className="w-full">
                <h2 className="text-xl font-bold mb-2">ライバルユーザーリストの追加と削除</h2>
            </div>
            {isLoading ? (
                <p>Loading...</p>
            ) : (
                <div className="overflow-x-auto">
                <table className="min-w-full border-collapse">
                    <thead>
                        <tr>
                            <th className="px-6 py-4 text-left">Rival Atcoder ID</th>
                        </tr>
                    </thead>
                    <tbody>
                        {rivals?.map((rival, index) => (
                            <EditingRivalItem
                                key={index}
                                atcoder_id={rival.atcoder_id}
                            />
                        ))}
                    </tbody>

                </table>
                <form onSubmit={submitRivalHandler}>
                    <input
                        className="mb-3 px-3 text-sm py-2 border border-gray-300"
                        name="id"
                        type="id"
                        placeholder="ライバルのAtCoder ID"
                        onChange={(e) => setNewId(e.target.value)}
                        value={newId}
                    />
                    <button
                        type="submit"
                        className="bg-blue-500 text-white py-2 px-4 rounded"
                    >
                        追加
                    </button>
                </form>
            </div>
            )}
        </div>
    )
}
