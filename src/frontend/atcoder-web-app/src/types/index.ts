export type Table = {
    id: string
    atcoder_id: string
    rating: number
    streak: number
}

export type Submission = {
    atcoder_id: string
    time: string
    problem: string
    diff: number
}

export type Rival = {
    id: string
    atcoder_id: string
}

export type CsrfToken = {
    csrf_token: string
}

export type Credential = {
    email: string
    password: string
}