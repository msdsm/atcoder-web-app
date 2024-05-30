import { FC } from 'react'
import { SubmissionItem } from './SubmissionItem'
import { useQuerySubmissions } from '../hooks/useQuerySubmissions'

export const SubmissionList: FC = () => {
    const {data, isLoading } = useQuerySubmissions()
    return (
        <div>
            {isLoading ? (
                <p>Loading...</p>
            ) : (
            <table>
                <thead>
                    <tr>
                        <th>Atcoder ID</th>
                        <th>Time</th>
                        <th>Problem</th>
                        <th>Diff</th>
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
            </table>)}
        </div>
    )
}