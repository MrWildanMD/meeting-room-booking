<script lang="ts">
  import { Button, Card, DatePicker, InputNumber, Select } from "stwui";
  import bg from "$lib/assets/img/bg.jpg";
  import bg1 from "$lib/assets/img/bg1.svg";
  import bg2 from "$lib/assets/img/bg2.svg";
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";
  import dayjs from "dayjs";
  import { checkIn, checkOut, totalGuest, additionalItems } from "$lib/stores/booking-store";
  import { id, name, privyId, typeUser, email, teleId } from "$lib/stores/user-store";

  import type { Items as ItemsModel } from "../../models/items";
  import type { SelectOption } from "stwui/types";
  import type { PageData } from "./$types";

  let guestNumber: number;
  let checkInDate: Date;
  let checkOutDate: Date;
  let addItems: Array<ItemsModel>;
  let today = dayjs();
  let yesterday = dayjs(today).subtract(1, "day").toDate();
  let isMobile: boolean;
  export let data: PageData;

  id.set(data.user.ID);
  name.set(data.user.name);
  privyId.set(data.user.privy_id);
  typeUser.set(data.user.type_user);
  email.set(data.user.email);
  teleId.set(data.user.tele_id);

  const items: SelectOption[] = [
    {
      value: "white_board",
      label: "White Board",
    },
    {
      value: "spidol",
      label: "Spidol",
    },
    {
      value: "converter",
      label: "Converter",
    },
    {
      value: "monitor",
      label: "Monitor",
    },
  ];

  $: bgImage = `background-image: url("${bg}");`;
  $: bgImage1 = `background-image: url("${bg1}");`;
  $: bgImage2 = `background-image: url("${bg2}");`;
  $: visible = false;
  $: roomVisible = false;
  $: inputGuestError =
    guestNumber == undefined || guestNumber == 0
      ? "Number of guest is required or not below 1"
      : "";
  $: checkInError = checkInDate == undefined ? "Check-in date is required" : "";
  $: checkOutError = checkOutDate == undefined ? "Check-out date is required" : "";

  onMount(() => {
    isMobile = window.innerWidth < 768;
    visible = true;
  });

  function scrollToFillForm(e: any) {
    e.preventDefault();
    document.getElementById("fill")!.scrollIntoView({ behavior: "smooth" });
  }
  function scrollToRoom(e: any) {
    e.preventDefault();
    let additional = addItems.map((items) => items.label).join(",");
    if (!inputGuestError && !checkInError && !checkOutError) {
      roomVisible = true;
      checkIn.set(dayjs(checkInDate).format());
      checkOut.set(dayjs(checkOutDate).format());
      totalGuest.set(guestNumber);
      additionalItems.set(additional);
      document.getElementById("room")!.scrollIntoView({ behavior: "smooth" });
    }
  }

  const additionalIcon =
    '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>playlist-edit</title><path d="M3 6V8H14V6H3M3 10V12H14V10H3M20 10.1C19.9 10.1 19.7 10.2 19.6 10.3L18.6 11.3L20.7 13.4L21.7 12.4C21.9 12.2 21.9 11.8 21.7 11.6L20.4 10.3C20.3 10.2 20.2 10.1 20 10.1M18.1 11.9L12 17.9V20H14.1L20.2 13.9L18.1 11.9M3 14V16H10V14H3Z" /></svg>';
  const calendarIcon =
    '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>calendar-clock</title><path d="M15,13H16.5V15.82L18.94,17.23L18.19,18.53L15,16.69V13M19,8H5V19H9.67C9.24,18.09 9,17.07 9,16A7,7 0 0,1 16,9C17.07,9 18.09,9.24 19,9.67V8M5,21C3.89,21 3,20.1 3,19V5C3,3.89 3.89,3 5,3H6V1H8V3H16V1H18V3H19A2,2 0 0,1 21,5V11.1C22.24,12.36 23,14.09 23,16A7,7 0 0,1 16,23C14.09,23 12.36,22.24 11.1,21H5M16,11.15A4.85,4.85 0 0,0 11.15,16C11.15,18.68 13.32,20.85 16,20.85A4.85,4.85 0 0,0 20.85,16C20.85,13.32 18.68,11.15 16,11.15Z" /></svg>';
  const guestIcon =
    '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>account-group</title><path d="M12,5.5A3.5,3.5 0 0,1 15.5,9A3.5,3.5 0 0,1 12,12.5A3.5,3.5 0 0,1 8.5,9A3.5,3.5 0 0,1 12,5.5M5,8C5.56,8 6.08,8.15 6.53,8.42C6.38,9.85 6.8,11.27 7.66,12.38C7.16,13.34 6.16,14 5,14A3,3 0 0,1 2,11A3,3 0 0,1 5,8M19,8A3,3 0 0,1 22,11A3,3 0 0,1 19,14C17.84,14 16.84,13.34 16.34,12.38C17.2,11.27 17.62,9.85 17.47,8.42C17.92,8.15 18.44,8 19,8M5.5,18.25C5.5,16.18 8.41,14.5 12,14.5C15.59,14.5 18.5,16.18 18.5,18.25V20H5.5V18.25M0,20V18.5C0,17.11 1.89,15.94 4.45,15.6C3.86,16.28 3.5,17.22 3.5,18.25V20H0M24,20H20.5V18.25C20.5,17.22 20.14,16.28 19.55,15.6C22.11,15.94 24,17.11 24,18.5V20Z" /></svg>';
</script>

{#if visible}
  <div transition:fade={{ delay: 250, duration: 500 }}>
    <main>
      <div
        class="hero flex flex-col justify-center items-center h-screen bg-local bg-center bg-cover"
        style={bgImage}
      >
        <div class="px-4 md:w-8/12">
          <p
            class="md:text-center text-4xl md:text-5xl font-bold mb-10 text-red-500 break-words text-stroke-white"
          >
            "Book Meeting Rooms with Ease. Simplify Your Workspace Planning.
            Productive Meetings Await."
          </p>
        </div>
        <Button type="danger" size="xl" on:click={scrollToFillForm}
          >Reserve Now!</Button
        >
      </div>
      <div
        id="fill"
        class="hero flex flex-col justify-center items-center h-screen bg-local bg-center bg-cover"
        style={bgImage1}
      >
        <Card bordered elevation="lg" class="w-9/12">
          <Card.Content>
            <h3 class="mb-4">Fill out this form</h3>
            <div class="flex flex-col gap-4">
              <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col flex-auto gap-2">
                  <DatePicker
                    name="check-in"
                    label="Check-In"
                    showTime
                    bind:value={checkInDate}
                    format="MMMM D, YYYY - H:mm"
                    placeholder="MMMM D, YYYY - H:mm"
                    mobile={isMobile}
                    error={checkInError}
                    min={yesterday}
                  >
                    <DatePicker.Label slot="label">Check-In</DatePicker.Label>
                    <DatePicker.Leading slot="leading" data={calendarIcon} />
                  </DatePicker>
                </div>
                <div class="flex flex-col flex-auto gap-2">
                  <DatePicker
                    name="check-out"
                    label="Check-Out"
                    showTime
                    format="MMMM D, YYYY - H:mm"
                    bind:value={checkOutDate}
                    placeholder="MMMM D, YYYY - H:mm"
                    mobile={isMobile}
                    error={checkOutError}
                    min={yesterday}
                  >
                    <DatePicker.Label slot="label">Check-Out</DatePicker.Label>
                    <DatePicker.Leading slot="leading" data={calendarIcon} />
                  </DatePicker>
                </div>
              </div>
              <div class="flex flex-col gap-2">
                <InputNumber
                  name="number-of-guest"
                  allowClear
                  bind:value={guestNumber}
                  placeholder="Total guest"
                  error={inputGuestError}
                >
                  <InputNumber.Label slot="label"
                    >Number of Guest</InputNumber.Label
                  >
                  <InputNumber.Leading slot="leading" data={guestIcon} />
                </InputNumber>
              </div>
              <div class="flex flex-col gap-2">
                <Select
                  mobile={isMobile}
                  name="select-4"
                  placeholder="Monitor, etc.."
                  multiple
                  bind:value={addItems}
                >
                  <Select.Label slot="label">Additional Items</Select.Label>
                  <Select.Leading slot="leading" data={additionalIcon} />
                  <Select.Options slot="options">
                    {#each items as option}
                      <Select.Options.Option {option} />
                    {/each}
                  </Select.Options>
                </Select>
              </div>
              <Button
                class="bg-red-500 hover:bg-red-600 text-white font-bold py-4 px-8 rounded"
                on:click={scrollToRoom}>Select Room</Button
              >
            </div>
          </Card.Content>
        </Card>
      </div>
      {#if roomVisible}
        <div
          id="room"
          class="hero flex flex-col justify-start items-center h-full bg-local bg-center bg-cover py-8 gap-8"
          style={bgImage2}
        >
          <h1>ROOMS</h1>
          <form method="POST" action="?/create">
            <div class="grid grid-cols-1 gap-4 md:grid-cols-4 px-4">
              <input type="hidden" name="user_id" value="{$id}">
              <input type="hidden" name="check_in" value="{$checkIn}">
              <input type="hidden" name="check_out" value="{$checkOut}">
              <input type="hidden" name="guest_total" value="{$totalGuest}">
              <input type="hidden" name="additional_items" value="{$additionalItems}">
              <input type="hidden" name="approval_id" value="{1}">
              {#each data.rooms as room}
              <input type="hidden" name="room_id" value="{room.ID}">
              <Card>
                <Card.Header class="font-bold" slot="header"
                  >{room.title}</Card.Header
                >
                <Card.Content slot="content">
                  <p class="card-text text-gray-700">
                    {room.description}
                  </p>
                  <p class="card-text text-gray-700">
                    Location: {room.location}
                  </p>
                  <p class="card-text text-gray-700">Type: {#if room.type === 1}
                     Kecil
                  {:else if  room.type === 2}
                     Sedang
                  {:else}
                     Besar
                  {/if}
                   </p>
                  <p class="card-text text-gray-700">Capacity: {room.capacity}</p>
                  <p class="card-text text-gray-700">Status: {#if room.status === 1}
                     Booked
                  {/if} Available </p>
                </Card.Content>
                <Card.Footer class="text-center" slot="footer">
                  <Button disabled={room.status == 1} class="w-full" htmlType="submit" type="danger"
                    >SELECT ROOM</Button
                  >
                </Card.Footer>
              </Card>
              {/each}
            </div>
          </form>
        </div>
      {/if}
    </main>
  </div>
{/if}

<style>
  .hero {
    border-radius: 0.25rem;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.125);
  }

  .text-stroke-white {
    text-shadow: 2px 2px 0 #fff, -1px -1px 0 #fff, 1px -1px 0 #fff,
      -1px 1px 0 #fff, 2px 2px 0 #fff;
  }
</style>
