import type { PageLoad } from "./$types"
import { PUBLIC_API_URL } from "$env/static/public"

export const load = (async ({ fetch }) => {
  try {
    const url = PUBLIC_API_URL

    const res = await fetch(url)
    const data = await res.json()

    if (data) {
      return { data }
    }

    return {}
  } catch (err) {
    return {}
  }
}) satisfies PageLoad
