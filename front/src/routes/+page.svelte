<script lang="ts">
  const lastUpdated = new Date();
  const isoString = lastUpdated.toLocaleString();

  interface ViolationInfo {
    _id: string;
    captureTime: string;
    closestDistance: number;
    pilot: {
      pilotId: string;
      firstName: string;
      lastName: string;
      phoneNumber: string;
      createdDt: string;
      email: string;
    };
  }

  export let data: {
    data: ViolationInfo[];
  };

  // Automatically refesh every 10 minutes
  setInterval(() => {
    location.reload();
  }, 10 * 60 * 1000);
</script>

<div class="flex flex-col gap-10">
  <div class="text-center">
    <h1 class="text-3xl text-sky-500 mb-4">Welcome to NDZ Violation List</h1>
    <p>Last updated: {isoString}</p>
  </div>
  {#if data.data}
    <table class="text-left text-gray-500">
      <thead class="text-gray-700 uppercase bg-gray-50">
        <tr>
          <th class="py-3 px-6">Name</th>
          <th class="py-3 px-6">Email</th>
          <th class="py-3 px-6">Closest distance (m)</th>
          <th class="py-3 px-6">Capture time</th>
        </tr>
      </thead>
      <tbody />
      {#each data.data as rec}
        <tr class="bg-white border-b">
          <td class="py-4 px-6 font-medium text-gray-900">
            {rec.pilot.firstName + " " + rec.pilot.lastName}
          </td>
          <td class="py-4 px-6">
            {rec.pilot.email}
          </td>
          <td class="py-4 px-6 text-right">
            {(rec.closestDistance / 1000).toFixed(2)}
          </td>
          <td class="py-4 px-6">
            {new Date(rec.captureTime).toLocaleString()}
          </td>
        </tr>
      {/each}
    </table>
  {/if}
</div>
