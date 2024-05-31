import { FC } from 'react'
import { SubmissionItem } from './SubmissionItem'
import { useQuerySubmissions } from '../hooks/useQuerySubmissions'

export const SubmissionList: FC = () => {
    const {data, isLoading } = useQuerySubmissions()
    return (
        <div>
            <div className="w-full">
                <h2 className="text-xl font-bold mb-2">今日の提出</h2>
            </div>
            {isLoading ? (
                <p>Loading...</p>
            ) : (
            <div className="overflow-x-auto">
                <table className="min-w-full border-collapse">
                    <thead>
                        <tr>
                            <th className="px-6 py-4 text-left">Atcoder ID</th>
                            <th className="px-6 py-4 text-left">Time</th>
                            <th className="px-6 py-4 text-left">Problem</th>
                            <th className="px-6 py-4 text-left">Diff</th>
                        </tr>
                    </thead>
                    <tbody>
                        {data?.map((submission, index) => (
                            <SubmissionItem
                                key={index}
                                atcoder_id={submission.atcoder_id}
                                time={submission.time}
                                problem={submission.problem}
                                diff={submission.diff}
                            />
                        ))}
                    </tbody>
                </table>
            </div>)}
        </div>
    )
}