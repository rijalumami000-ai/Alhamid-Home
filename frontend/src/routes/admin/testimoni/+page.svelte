<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Testimonial } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Star } from 'lucide-svelte';

  let list = $state<Testimonial[]>([]);
  let isModalOpen = $state(false);
  let editing = $state<Testimonial | null>(null);
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

  async function handleDelete(id: number) {
    if (confirm('Hapus testimoni ini?')) {
      try {
        await api.admin.deleteTestimonial(id);
        await loadData();
      } catch (err: any) {
        alert(err.message || 'Gagal menghapus');
      }
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
      <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
        Testimoni Alumni & Wali Santri
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola ulasan dan kutipan pengalaman belajar di Al-Hamid.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white text-xs font-bold shadow-md transition"
    >
      <Plus class="w-4 h-4" /> Tambah Testimoni
    </button>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-200 p-6 flex flex-col justify-between shadow-sm space-y-4">
        <div class="space-y-3">
          <div class="flex items-center gap-1 text-amber-400">
            {#each Array(item.rating || 5) as _}
              <Star class="w-4 h-4 fill-amber-400" />
            {/each}
          </div>
          <p class="text-slate-600 text-xs sm:text-sm leading-relaxed italic">
            "{item.content}"
          </p>
        </div>

        <div class="pt-3 border-t border-slate-100 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <img src={item.avatar_url} alt={item.name} class="w-10 h-10 rounded-full object-cover border border-emerald-600" />
            <div>
              <div class="font-bold text-xs text-slate-900">{item.name}</div>
              <div class="text-[11px] text-slate-400">{item.role}</div>
            </div>
          </div>
          <div class="space-x-1">
            <button onclick={() => openEdit(item)} class="p-1.5 rounded-lg text-slate-500 hover:text-emerald-700">
              <Edit2 class="w-4 h-4" />
            </button>
            <button onclick={() => handleDelete(item.id)} class="p-1.5 rounded-lg text-slate-500 hover:text-red-600">
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
        <h2 class="font-serif-luxury font-bold text-lg text-slate-900">
          {editing ? 'Edit Testimoni' : 'Tambah Testimoni'}
        </h2>
        <button onclick={() => (isModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSave} class="p-6 space-y-4">
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

        <div class="space-y-1">
          <label for="t_avatar" class="text-xs font-bold text-slate-700 uppercase">URL Foto Avatar</label>
          <input
            id="t_avatar"
            type="url"
            bind:value={form.avatar_url}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

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
            class="px-5 py-2 rounded-xl bg-emerald-800 text-white text-xs font-bold shadow-md"
          >
            Simpan Testimoni
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

