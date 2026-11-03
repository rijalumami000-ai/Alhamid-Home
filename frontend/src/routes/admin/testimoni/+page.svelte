<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Testimonial } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Star } from 'lucide-svelte';
  import ConfirmDeleteModal from '$lib/components/ConfirmDeleteModal.svelte';
  import DriveImageInput from '$lib/components/DriveImageInput.svelte';

  let list = $state<Testimonial[]>([]);
  let isModalOpen = $state(false);
  let editing = $state<Testimonial | null>(null);

  // Delete modal state
  let isDeleteModalOpen = $state(false);
  let testimonialToDelete = $state<Testimonial | null>(null);
  let isDeleting = $state(false);

  let form = $state({
    name: '',
    role: '',
    content: '',
    avatar_url: '',
    rating: 5,
    order: 0,
    is_active: true,
  });

  async function loadData() {
    try {
      list = await api.admin.getTestimonials();
    } catch (e) {
      console.error(e);
    }
  }

  function openCreate() {
    editing = null;
    form = {
      name: '',
      role: 'Alumni / Wali Santri',
      content: '',
      avatar_url: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=400&q=80',
      rating: 5,
      order: list.length + 1,
      is_active: true,
    };
    isModalOpen = true;
  }

  function openEdit(t: Testimonial) {
    editing = t;
    form = {
      name: t.name,
      role: t.role,
      content: t.content,
      avatar_url: t.avatar_url,
      rating: t.rating,
      order: t.order,
      is_active: t.is_active,
    };
    isModalOpen = true;
  }

  async function handleSave(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editing) {
        await api.admin.updateTestimonial(editing.id, form);
      } else {
        await api.admin.createTestimonial(form);
      }
      isModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan testimoni');
    }
  }

  function promptDelete(t: Testimonial) {
    testimonialToDelete = t;
    isDeleteModalOpen = true;
  }

  async function confirmDelete() {
    if (!testimonialToDelete) return;
    isDeleting = true;
    try {
      await api.admin.deleteTestimonial(testimonialToDelete.id);
      isDeleteModalOpen = false;
      testimonialToDelete = null;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menghapus testimoni');
    } finally {
      isDeleting = false;
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kelola Testimoni | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
        Testimoni Alumni & Wali Santri
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola ulasan dan kutipan pengalaman belajar di Al-Hamid.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition active:scale-95"
    >
      <Plus class="w-4 h-4" /> Tambah Testimoni
    </button>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm flex flex-col justify-between space-y-4">
        <div class="space-y-3">
          <div class="flex items-center gap-1 text-amber-400">
            {#each Array(item.rating || 5) as _}
              <Star class="w-4 h-4 fill-amber-400 text-amber-400" />
            {/each}
          </div>
          <p class="text-xs text-slate-600 italic line-clamp-4 leading-relaxed">
            "{item.content}"
          </p>
        </div>

        <div class="pt-4 border-t border-slate-100 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-full bg-slate-200 overflow-hidden shrink-0">
              <img src={item.avatar_url} alt={item.name} class="w-full h-full object-cover" />
            </div>
            <div>
              <div class="font-bold text-xs text-slate-900">{item.name}</div>
              <div class="text-[10px] text-slate-400">{item.role}</div>
            </div>
          </div>
          <div class="space-x-1">
            <button
              onclick={() => openEdit(item)}
              class="p-1.5 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-slate-100 transition"
              title="Edit Testimoni"
            >
              <Edit2 class="w-4 h-4" />
            </button>
            <button
              onclick={() => promptDelete(item)}
              class="p-1.5 rounded-lg text-slate-500 hover:text-rose-600 hover:bg-rose-50 transition"
              title="Hapus Testimoni"
            >
              <Trash2 class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>

<!-- MODAL FORM -->
{#if isModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-md rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-display font-bold text-lg text-slate-900">
          {editing ? 'Edit Testimoni' : 'Tambah Testimoni'}
        </h2>
        <button onclick={() => (isModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSave} class="p-6 space-y-4 max-h-[80vh] overflow-y-auto">
        <div class="space-y-1">
          <label for="t_name" class="text-xs font-bold text-slate-700 uppercase">Nama Lengkap *</label>
          <input
            id="t_name"
            type="text"
            required
            bind:value={form.name}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="space-y-1">
          <label for="t_role" class="text-xs font-bold text-slate-700 uppercase">Status / Peran *</label>
          <input
            id="t_role"
            type="text"
            required
            bind:value={form.role}
            placeholder="Alumni 2020 / Wali Santri MA Cintamulya"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <DriveImageInput
          id="t_avatar"
          label="Foto Avatar / Profil"
          bind:value={form.avatar_url}
          placeholder="https://drive.google.com/... atau https://..."
        />

        <div class="space-y-1">
          <label for="t_content" class="text-xs font-bold text-slate-700 uppercase">Isi Kutipan / Testimoni *</label>
          <textarea
            id="t_content"
            required
            rows="4"
            bind:value={form.content}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>

        <div class="pt-4 border-t border-slate-100 flex justify-end gap-2">
          <button
            type="button"
            onclick={() => (isModalOpen = false)}
            class="px-4 py-2 rounded-xl border border-slate-200 text-xs font-bold"
          >
            Batal
          </button>
          <button
            type="submit"
            class="px-5 py-2 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition"
          >
            Simpan Testimoni
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- CONFIRM DELETE MODAL -->
<ConfirmDeleteModal
  isOpen={isDeleteModalOpen}
  title="Hapus Testimoni"
  itemName={testimonialToDelete ? `${testimonialToDelete.name} (${testimonialToDelete.role})` : ''}
  message="Kutipan testimoni ini akan dihapus dari daftar ulasan beranda dan profil."
  {isDeleting}
  onConfirm={confirmDelete}
  onCancel={() => {
    isDeleteModalOpen = false;
    testimonialToDelete = null;
  }}
/>
