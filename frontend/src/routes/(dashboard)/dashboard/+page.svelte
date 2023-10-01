<script lang="ts">
  import dayjs from "dayjs";
  import type { ActionData, PageData } from "./$types";
  import type { Booking } from "../../../models/booking";
  import {
    Button,
    Input,
    InputNumber,
    Modal,
    Portal,
    Select,
    Tabs,
    TextArea,
  } from "stwui";
  import type { Room } from "../../../models/rooms";
  import type { SelectOption } from "stwui/types";

  let selectedBooking: Booking;
  let selectedRoom: Room;
  let modalBookingOpen = false;
  let modalRoomOpen = false;
  let addRoom = false;
  let bookings: Booking[];
  let rooms: Room[];

  $: {
    bookings = data.bookings;
    rooms = data.rooms;
  }

  function showBookingDetails(booking: Booking) {
    selectedBooking = booking;
    modalBookingOpen = true;
  }

  function showRoomDetails(room: Room) {
    selectedRoom = room;
    modalRoomOpen = true;
  }

  interface Tab {
    href: string;
    title: string;
  }

  const tabs: Tab[] = [
    {
      href: "#tab1",
      title: "Bookings",
    },
    {
      href: "#tab2",
      title: "Rooms",
    },
  ];

  const roomType: SelectOption[] = [
    {
      value: 1,
      label: "Kecil",
    },
    {
      value: 2,
      label: "Sedang",
    },
    {
      value: 3,
      label: "Besar",
    },
  ];
  const roomStatus: SelectOption[] = [
    {
      value: 0,
      label: "Available",
    },
    {
      value: 1,
      label: "Booked",
    },
  ];

  let currentTab = "#tab1";
  let currentIndex = 0;

  export let data: PageData;
  export let form: ActionData;
</script>

<main>
  <div class="container">
    <Tabs {currentTab}>
      {#each tabs as tab, i}
        <Tabs.Tab
          key={tab.href}
          href={tab.href}
          on:click={() => {
            (currentTab = tab.href), (currentIndex = i);
          }}
        >
          {tab.title}
        </Tabs.Tab>
      {/each}
    </Tabs>
    {#if currentIndex == 0}
      <!-- Booking Data -->
      <div class="overflow-x-auto">
        <table class="daisytable">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Check-in</th>
              <th>Check-out</th>
              <th>Guest</th>
              <th class="text-center">Status</th>
              <th>Additional Items</th>
              <th>Created At</th>
              <th class="text-center">Action</th>
            </tr>
          </thead>
          <tbody>
            {#each bookings as booking}
              <tr>
                <td>{booking.ID}</td>
                <td>{booking.user.name}</td>
                <td>{dayjs(booking.check_in).format("YYYY-MM-DD HH:mm")}</td>
                <td>{dayjs(booking.check_out).format("YYYY-MM-DD HH:mm")}</td>
                <td>{booking.number_of_guests}</td>
                <td>
                  <span
                    class="px-2 py-1 {booking.booking_status == 0
                      ? 'bg-orange-200 text-orange-800'
                      : ''} bg-green-200 text-green-800 rounded-full"
                    >{#if booking.booking_status == 1}
                      Approved
                    {:else}
                      Pending
                    {/if}</span
                  >
                </td>
                <td>{booking.additional_item}</td>
                <td>{dayjs(booking.CreatedAt).format("YYYY-MM-DD HH:mm")}</td>
                <th>
                  <button
                    on:click={() => showBookingDetails(booking)}
                    class="daisybtn daisybtn-ghost daisybtn-xs">Details</button
                  >
                </th>
              </tr>
            {/each}
          </tbody>
          <tfoot>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Check-in</th>
              <th>Check-out</th>
              <th>Guest</th>
              <th class="text-center">Status</th>
              <th>Additional Items</th>
              <th>Created At</th>
              <th class="text-center">Action</th>
            </tr>
          </tfoot>
        </table>
      </div>
    {:else}
      <!-- Rooms Data -->
      <Button
        on:click={() => {
          (modalRoomOpen = true), (addRoom = true);
        }}
        class="bg-green-500 hover:bg-green-600 text-white my-4 mx-4"
      >
        Add Room
      </Button>
      <div class="overflow-x-auto">
        <table class="daisytable">
          <thead>
            <tr>
              <th>ID</th>
              <th>Title</th>
              <th>Description</th>
              <th>Location</th>
              <th>Type</th>
              <th>Capacity</th>
              <th>Status</th>
              <th class="text-center">Action</th>
            </tr>
          </thead>
          <tbody>
            {#each rooms as room}
              <tr>
                <td>{room.ID}</td>
                <td>{room.title}</td>
                <td>{room.description}</td>
                <td>{room.location}</td>
                <td
                  >{#if room.type === 1}
                    Kecil
                  {:else if room.type === 2}
                    Sedang
                  {:else}
                    Besar
                  {/if}</td
                >
                <td>
                  {room.capacity}
                </td>
                <td
                  ><span
                    class="px-2 py-1 {room.status == 1
                      ? 'bg-orange-200 text-orange-800'
                      : ''} bg-green-200 text-green-800 rounded-full"
                    >{#if room.status == 1}
                      Booked
                    {:else}
                      Open
                    {/if}</span
                  ></td
                >
                <th>
                  <button
                    on:click={() => showRoomDetails(room)}
                    class="daisybtn daisybtn-ghost daisybtn-xs">Details</button
                  >
                </th>
              </tr>
            {/each}
          </tbody>
          <tfoot>
            <tr>
              <th>ID</th>
              <th>Title</th>
              <th>Description</th>
              <th>Location</th>
              <th>Type</th>
              <th>Capacity</th>
              <th>Status</th>
              <th class="text-center">Action</th>
            </tr>
          </tfoot>
        </table>
      </div>
    {/if}
  </div>
  <Portal>
    {#if modalBookingOpen}
      <Modal handleClose={() => (modalBookingOpen = false)}>
        <Modal.Content slot="content">
          <Modal.Content.Header class="font-bold" slot="header"
            >BOOKING SUMMARY</Modal.Content.Header
          >
          <Modal.Content.Body slot="body">
            <div class="grid grid-cols-2 gap-4">
              <div>
                <h3 class="text-gray-600">Check-In:</h3>
                <p class="text-black">
                  {dayjs(selectedBooking.check_in).format("YYYY-MM-DD HH:mm")}
                </p>
              </div>
              <div>
                <h3 class="text-gray-600">Check-Out:</h3>
                <p class="text-black">
                  {dayjs(selectedBooking.check_out).format("YYYY-MM-DD HH:mm")}
                </p>
              </div>
              <div>
                <h3 class="text-gray-600">Created At:</h3>
                <p class="text-black">
                  {dayjs(selectedBooking.CreatedAt).format("YYYY-MM-DD HH:mm")}
                </p>
              </div>
              <div>
                <h3 class="text-gray-600">Number of Guests:</h3>
                <p class="text-black">
                  {selectedBooking.number_of_guests}
                </p>
              </div>
              <div>
                <h3 class="text-gray-600">Booking Status:</h3>
                <span
                  class="px-2 py-1 {selectedBooking.booking_status == 0
                    ? 'bg-orange-200 text-orange-800'
                    : ''} bg-green-200 text-green-800 rounded-full"
                  >{#if selectedBooking.booking_status == 1}
                    Approved
                  {:else}
                    Pending
                  {/if}</span
                >
              </div>
              <div>
                <h3 class="text-gray-600">Additional Items:</h3>
                <p class="text-black">
                  {selectedBooking.additional_item}
                </p>
              </div>
              <div>
                <h3 class="text-gray-600">User Name:</h3>
                <p class="text-black">{selectedBooking.user.name}</p>
              </div>
              <div>
                <h3 class="text-gray-600">Room Title:</h3>
                <p class="text-black">{selectedBooking.room.title}</p>
              </div>
            </div>
            <div class="flex flex-row justify-between pt-4">
              <form action="?/approve" method="post">
                {#if form}
                  <h4 class="text-red-600 font-light text-md text-center">
                    {form.error}
                  </h4>
                {/if}
                <input
                  name="booking_id"
                  type="hidden"
                  value={selectedBooking.ID}
                />
                <input
                  name="room_id"
                  type="hidden"
                  value={selectedBooking.room.ID}
                />
                <Button
                  htmlType="submit"
                  type="danger"
                  class="bg-green-500 hover:bg-green-600 text-white font-bold rounded"
                  >Approve</Button
                >
              </form>
              <form action="?/decline" method="post">
                {#if form}
                  <h4 class="text-red-600 font-light text-md text-center">
                    {form.error}
                  </h4>
                {/if}
                <input
                  name="booking_id"
                  type="hidden"
                  value={selectedBooking.ID}
                />
                <Button
                  htmlType="submit"
                  on:SubmitEvent={(e) => e.preventDefault()}
                  type="danger"
                  class="bg-red-500 hover:bg-red-600 text-white font-bold rounded"
                  >Decline</Button
                >
              </form>
            </div>
          </Modal.Content.Body>
        </Modal.Content>
      </Modal>
    {/if}
    {#if modalRoomOpen}
      <Modal handleClose={() => ((modalRoomOpen = false), (addRoom = false))}>
        <Modal.Content slot="content">
          <Modal.Content.Header class="font-bold" slot="header"
            >{#if addRoom}
              Add Room
            {:else}
              Room Details
            {/if}</Modal.Content.Header
          >
          <Modal.Content.Body slot="body">
            <form action="?/{addRoom ? 'create' : 'update'}" method="post">
              {#if form}
                <h4 class="text-red-600 font-light text-md text-center">
                  {form.error}
                </h4>
              {/if}
              {#if !addRoom}
                <input type="hidden" name="id" value={selectedRoom.ID} />
              {/if}
              <div class="mb-4">
                <Input name="title" value={addRoom ? "" : selectedRoom.title}>
                  <Input.Label slot="label">Title</Input.Label>
                </Input>
              </div>
              <div class="mb-4">
                <TextArea
                  value={addRoom ? "" : selectedRoom.description}
                  name="description"
                  placeholder="Input description"
                >
                  <TextArea.Label slot="label">Description</TextArea.Label>
                </TextArea>
              </div>
              <div class="mb-4">
                <Input
                  name="location"
                  value={addRoom ? "" : selectedRoom.location}
                >
                  <Input.Label slot="label">Location</Input.Label>
                </Input>
              </div>
              <div class="mb-4">
                <Select name="type">
                  <Select.Label slot="label">Room Type</Select.Label>
                  <Select.Options slot="options">
                    {#each roomType as option}
                      <Select.Options.Option {option} />
                    {/each}
                  </Select.Options>
                </Select>
              </div>
              <div class="mb-4">
                <InputNumber
                  name="capacity"
                  value={addRoom ? 0 : selectedRoom.capacity}
                >
                  <InputNumber.Label slot="label">Capacity</InputNumber.Label>
                </InputNumber>
              </div>
              <!-- {#if !addRoom}
                 <div class="mb-4">
                   <Select name="status">
                     <Select.Label slot="label">Room Status</Select.Label>
                     <Select.Options slot="options">
                       {#each roomStatus as option}
                         <Select.Options.Option {option} />
                       {/each}
                     </Select.Options>
                   </Select>
                 </div>
              {/if} -->
              <div class="mt-6">
                <Button
                  htmlType="submit"
                  on:submit={(e) => e.preventDefault()}
                  class="bg-blue-500 text-white px-4 py-2 rounded-md w-full hover:bg-blue-600"
                  >{#if addRoom}
                    Submit
                  {:else}
                    Update
                  {/if}</Button
                >
              </div>
            </form>
            {#if !addRoom}
              <form action="?/delete" method="post">
                {#if form}
                  <h4 class="text-red-600 font-light text-md text-center">
                    {form.error}
                  </h4>
                {/if}
                <input name="room_id" type="hidden" value={selectedRoom.ID} />
                <Button
                  htmlType="submit"
                  on:submit={(e) => e.preventDefault()}
                  class="bg-red-500 text-white px-4 py-2 rounded-md w-full mt-4 hover:bg-red-600"
                  >Delete</Button
                >
              </form>
            {/if}
          </Modal.Content.Body>
        </Modal.Content>
      </Modal>
    {/if}
  </Portal>
</main>
