import { FC, FormEvent } from 'react'
import { TableItem } from './TableItem'
import { useQueryTables } from '../hooks/useQueryTables'
import useStore from '../store'

export const TableList: FC = () => {
    const {data, isLoading } = useQueryTables()
    const { editedRival } = useStore()
    return (
        <div>
            <div className="w-full">
                <h2 className="text-xl font-bold mb-2">ライバルユーザーリスト</h2>
            </div>
            {isLoading ? (
                <p>Loading...</p>
            ) : (
            <div className="overflow-x-auto">
                <table className="min-w-full border-collapse">
                    <thead>
                        <tr>
                            <th className="px-6 py-4 text-left">Atcoder ID</th>
                            <th className="px-6 py-4 text-left">Rating</th>
                            <th className="px-6 py-4 text-left">Streak</th>
                        </tr>
                    </thead>
                    <tbody>
                        {data?.map((table) => (
                            <TableItem
                                key={table.id}
                                id={table.id}
                                atcoder_id={table.atcoder_id}
                                rating={table.rating}
                                streak={table.streak}
                            />
                        ))}
                    </tbody>
                </table>
            </div>)}
        </div>
    )
}